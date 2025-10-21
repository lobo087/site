import json
from flask import Flask, request, jsonify

app = Flask(__name__)

# Inventario en memoria (se borra al reiniciar la app)
inventario = [
    {"titulo": "Cien años de soledad", "autor": "Gabriel García Márquez", "cantidad": 3},
    {"titulo": "Don Quijote de la Mancha", "autor": "Miguel de Cervantes", "cantidad": 2},
    {"titulo": "El Principito", "autor": "Antoine de Saint-Exupéry", "cantidad": 5},
    {"titulo": "1984", "autor": "George Orwell", "cantidad": 4},
    {"titulo": "Rayuela", "autor": "Julio Cortázar", "cantidad": 1}
]

@app.route("/libros", methods=["GET"])
def listar_libros():
    return jsonify(inventario), 200

@app.route("/libros", methods=["POST"])
def agregar_libro():
    # Espera JSON con al menos: titulo, autor
    data = request.get_json(silent=True) or {}
    titulo = data.get("titulo")
    autor = data.get("autor")
    cantidad = data.get("cantidad", 1)

    # Validaciones súper básicas
    if not titulo or not autor:
        return jsonify({"error": "Faltan campos requeridos: 'titulo' y 'autor'."}), 400
    try:
        cantidad = int(cantidad)
        if cantidad < 1:
            raise ValueError
    except (ValueError, TypeError):
        return jsonify({"error": "'cantidad' debe ser un entero >= 1."}), 400

    # Agregar al inventario
    libro = {"titulo": titulo, "autor": autor, "cantidad": cantidad}
    inventario.append(libro)
    return jsonify({"mensaje": "Libro agregado", "libro": libro}), 201

if __name__ == "__main__":
    app.run(debug=True)
