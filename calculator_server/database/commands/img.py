from typing import Optional
from database.commands.methods import available_methods
from database.migrations.files.img import CFileImg


def load_img(file, name=None) -> Optional[str]:
    img = CFileImg(
        filename=(name or file.filename),
        img_bytestream=file.stream
    )

    success = img.save(local=True)

    return img.name if success else None


def return_img(name) -> str:
    img = CFileImg(filename=name, from_db=True)
    img.show()

    return img.name


def process_img(name: str, methods: str, new_name) -> Optional[str]:
    img = CFileImg(filename=name, from_db=True)
    img.show()

    for method in methods:
        img = available_methods[method](img)

    img.recreate(new_name)
    success = img.save(local=True)
    if not success:
        return None

    img.show()
    return img.name

