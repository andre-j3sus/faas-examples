from flask import Flask, request, jsonify
import os

app = Flask(__name__)

@app.route("/write", methods=["POST"])
def write_file():
    """Writes content to a specified file (overwrites existing content)."""
    data = request.json
    path = data.get("path")
    content = data.get("content", "")

    if not path:
        return jsonify({"error": "Path is required"}), 400

    try:
        with open(path, "w") as f:
            f.write(content)
        return jsonify({"message": f"Content written to {path} successfully."})
    except ValueError as e:
        return jsonify({"error": str(e)}), 403

@app.route("/read", methods=["GET"])
def read_file():
    """Reads the content of a specified file."""
    path = request.args.get("path")

    if not path:
        return jsonify({"error": "Path is required"}), 400

    try:
        if not os.path.exists(path):
            return jsonify({"error": "File does not exist"}), 404

        with open(path, "r") as f:
            content = f.read()
        return jsonify({"message": f"Content of {path} read successfully.", "content": content})
    except ValueError as e:
        return jsonify({"error": str(e)}), 403

@app.route("/list", methods=["GET"])
def list_files():
    """Lists all files in the selected path."""
    
    path = request.args.get("path")

    if not path:
        return jsonify({"error": "Path is required"}), 400
    
    files = os.listdir(path)
    return jsonify({"files": files})

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)
