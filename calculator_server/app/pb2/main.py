from concurrent import futures
import grpc
import pb.base_pb2 as pb2
import pb.base_pb2_grpc as pb2_grpc


class Greeter(pb2_grpc.GreeterServicer):

    def SayHello(self, request, context):
        return pb2.HelloReply(message='Hello, %s!' % request.name)


class Meta(pb2_grpc.MetaServicer):

    def Blur(self, request, context):
        return pb2.BlurReply(name=img.name)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pb2_grpc.add_GreeterServicer_to_server(Greeter(), server)
    pb2_grpc.add_MetaServicer_to_server(Meta(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()

