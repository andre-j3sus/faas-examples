from flask import Flask, jsonify

app = Flask(__name__)

@app.route('/', methods=['GET'])
def hello_world():
    """Simple function that returns a greeting message."""
    return jsonify({
        "message": "Hello, World!",
        "status": "success"
    })

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=8080)