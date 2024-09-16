# Use the official Python image
FROM python:3.12-alpine

# Set the working directory
WORKDIR /app

# Copy the requirements file and install dependencies
COPY python_app/requirements.txt .

# Install dependencies, including PyTorch and transformers
RUN pip install --no-cache-dir -r requirements.txt

# Copy the application code and the fine-tuned model
COPY python_app/ .
COPY fine-tuned-model/ /app/fine-tuned-model

# Expose the gRPC port
EXPOSE 50052

# Run the Python application
CMD ["python", "app.py"]
