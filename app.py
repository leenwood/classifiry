from concurrent import futures
import grpc
from config.remote import model_pb2, model_pb2_grpc


# Подключение и инициализация модели классификации
# Пример использования простой модели, которая возвращает предопределенный тег

class Classifier(model_pb2_grpc.ClassifierServicer):
    def Classify(self, request, context):
        description = request.description
        # Логика классификации описания
        tag = self.classify_description(description)
        return model_pb2.ClassifyResponse(tag=tag)

    def classify_description(self, description):
        # Пример логики классификации. Здесь должен быть вызов вашей модели.
        if "example" in description:
            return "example_tag"
        return "default_tag"

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    model_pb2_grpc.add_ClassifierServicer_to_server(Classifier(), server)
    server.add_insecure_port('[::]:50051')
    print("Python gRPC server started on port 50051")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    serve()