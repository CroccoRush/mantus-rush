from database.migrations.files.file import DCFileBase
from database.migrations.files.img import DCFileImg
from database.migrations.meta.blur import DCMetaBlur
from database.migrations.meta.meta import DCMetaBase


def init():
    DCFileBase.SQL_create()
    DCFileImg.SQL_create()
    DCMetaBase.SQL_create()
    DCMetaBlur.SQL_create()


def start():
    init()
