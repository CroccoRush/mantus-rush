import PIL.Image
import PIL.ImageFilter

from database.migrations.files.img import CFileImg
from database.migrations.meta.blur import CMetaBlur


def img_blur(img: CFileImg) -> CFileImg:
    img.image = img.image.filter(filter=PIL.ImageFilter.BLUR())
    meta = CMetaBlur(img.id, exist=True)
    meta.save()
    img.add_meta(meta.table_name(), meta.id)
    return img
