from flask import Flask, request, jsonify, send_file
from PIL import Image
import io
import os
import tempfile

app = Flask(__name__)

@app.route('/resize', methods=['POST'])
def resize_image():
    """
    Resizes an image to specified width and height.
    
    Expects form data with:
    - 'image': The image file
    - 'width': Target width (integer)
    - 'height': Target height (integer)
    
    Returns the resized image.
    """
    if 'image' not in request.files:
        return jsonify({"error": "No image provided"}), 400
    
    try:
        width = int(request.form.get('width', 100))
        height = int(request.form.get('height', 100))
    except ValueError:
        return jsonify({"error": "Width and height must be integers"}), 400
    
    # Get the uploaded image
    image_file = request.files['image']
    
    try:
        # Open the image using PIL
        img = Image.open(image_file.stream)
        
        # Resize the image
        resized_img = img.resize((width, height))
        
        # Save to a temporary file
        temp_file = tempfile.NamedTemporaryFile(delete=False, suffix='.png')
        resized_img.save(temp_file.name)
        
        # Return the resized image
        return send_file(temp_file.name, mimetype='image/png', as_attachment=True, 
                        download_name=f"resized_{width}x{height}.png")
    
    except Exception as e:
        return jsonify({"error": str(e)}), 500

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=8080)