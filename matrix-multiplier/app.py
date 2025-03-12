from flask import Flask, request, jsonify
import numpy as np

app = Flask(__name__)

@app.route('/multiply', methods=['POST'])
def multiply_matrices():
    """
    Multiplies two matrices.
    
    Expects JSON with:
    - 'matrix_a': First matrix as a list of lists
    - 'matrix_b': Second matrix as a list of lists
    
    Returns the matrix product.
    """
    data = request.json
    
    if not data or 'matrix_a' not in data or 'matrix_b' not in data:
        return jsonify({"error": "Both matrices are required"}), 400
    
    try:
        # Convert input to numpy arrays
        matrix_a = np.array(data['matrix_a'], dtype=float)
        matrix_b = np.array(data['matrix_b'], dtype=float)
        
        # Check if matrices can be multiplied
        if matrix_a.shape[1] != matrix_b.shape[0]:
            return jsonify({
                "error": f"Cannot multiply matrices of shapes {matrix_a.shape} and {matrix_b.shape}. "
                         f"The number of columns in first matrix ({matrix_a.shape[1]}) must equal "
                         f"the number of rows in second matrix ({matrix_b.shape[0]})."
            }), 400
        
        # Perform matrix multiplication
        result = np.matmul(matrix_a, matrix_b)
        
        return jsonify({
            "result": result.tolist(),
            "shape": {
                "rows": result.shape[0],
                "columns": result.shape[1]
            }
        })
        
    except ValueError as e:
        return jsonify({"error": f"Invalid matrix format: {str(e)}"}), 400
    except Exception as e:
        return jsonify({"error": str(e)}), 500

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=8080)