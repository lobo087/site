# === Etapa base ===
FROM python:3.12-slim

# Define el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiamos los archivos del proyecto
COPY . .

# Instalar dependencias del archivo requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

# Exponemos el puerto (ajústalo si usas otro)
EXPOSE 5000

# (Opcional) Healthcheck para verificar que la API responde
HEALTHCHECK --interval=30s --timeout=3s --retries=3 \
  CMD wget -qO- http://127.0.0.1:5000/libros || exit 1

# Ejecutar con waitress (más estable que el servidor de desarrollo)
CMD ["waitress-serve", "--listen=0.0.0.0:5000", "app:app"]