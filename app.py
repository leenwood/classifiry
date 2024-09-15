from concurrent import futures
import grpc
from config.remote import model_pb2, model_pb2_grpc
import torch
from transformers import AutoTokenizer, AutoModelForSequenceClassification


# Загрузка маппера тегов
type_mapper = {
    'Негативный': 0,
    'Отравление': 1,
    'Позитивный': 2,
    'Проблемы с доставкой': 3,
    'Проблемы с едой': 4
}

# Класс для работы с моделью
class ModelHandler:
    def __init__(self, model_path):
        # Инициализация модели и токенайзера
        self.tokenizer = AutoTokenizer.from_pretrained(model_path)
        self.model = AutoModelForSequenceClassification.from_pretrained(model_path)
        self.model.eval()  # Переводим модель в режим оценки

    def predict(self, description):
        # Преобразование описания в формат, подходящий для модели
        inputs = self.tokenizer(description, return_tensors='pt', padding=True, truncation=True)
        # Предсказание модели
        with torch.no_grad():  # Отключаем градиенты, чтобы ускорить вычисления
            outputs = self.model(**inputs)
        # Определение метки
        predicted_class = torch.argmax(outputs.logits, dim=1).item()

        # Поиск текста метки по индексу
        for key, value in type_mapper.items():
            if value == predicted_class:
                return key
        return "Неизвестный тег"

# Подключение и инициализация модели классификации
# Пример использования простой модели, которая возвращает предопределенный тег

class Classifier(model_pb2_grpc.ClassifierServicer):
    def __init__(self, model_handler):
        # Используем переданный экземпляр модели
        self.model_handler = model_handler

    def Classify(self, request, context):
        description = request.description
        # Логика классификации описания
        tag = self.model_handler.predict(description)
        return model_pb2.ClassifyResponse(tag=tag)


def init_model(model_path):
    return ModelHandler(model_path)

def serve(model_handler, host: str):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    # Передаем единственный экземпляр ModelHandler в сервис Classifier
    model_servicer = Classifier(model_handler)
    model_pb2_grpc.add_ClassifierServicer_to_server(model_servicer, server)
    server.add_insecure_port(host)
    print(f"Python gRPC server started on port {host}")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    model = init_model("fine-tuned-model")
    print("model initialized")
    serve(model, '[::]:50051')