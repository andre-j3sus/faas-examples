from flask import Flask, request, jsonify
import numpy as np

app = Flask(__name__)

@app.route('/magnitude', methods=['POST'])
def calculate_magnitude():
    """
    Calculates the magnitude (Euclidean norm) of a vector.
    
    Expects JSON with:
    - 'vector': A list of numerical values representing a vector
    
    Returns the magnitude of the vector.
    """
    data = request.json
    
    if not data or 'vector' not in data:
        return jsonify({"error": "Vector is required"}), 400
    
    try:
        # Convert input to numpy array
        vector = np.array(data['vector'], dtype=float)
        
        # Calculate magnitude (Euclidean norm)
        magnitude = np.linalg.norm(vector)
        
        return jsonify({
            "vector": vector.tolist(),
            "magnitude": float(magnitude),
            "dimensions": len(vector)
        })
        
    except ValueError as e:
        return jsonify({"error": f"Invalid vector format: {str(e)}"}), 400
    except Exception as e:
        return jsonify({"error": str(e)}), 500

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=8080)