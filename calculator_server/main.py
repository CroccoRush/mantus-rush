from app.rest import api
from app.pb2 import main as pb


def main():
    api.start()
    pb.serve()


if __name__ == '__main__':
    main()
