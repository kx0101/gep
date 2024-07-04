from pptx import Presentation
from pptx.util import Inches
from datetime import datetime, timedelta
from pptx.oxml import parse_xml

def create_presentation():
    prs = Presentation()

    today = datetime.now()
    tomorrow = today + timedelta(days=1)

    screenshots = [
        f'images/{today.strftime("%d-%m-%Y")}.png',
        f'images/{tomorrow.strftime("%d-%m-%Y")}.png',
        f'images/{today.strftime("%d-%m-%Y")}-fire.jpg',
        "images/4.png",
        "images/5.png",
        "images/6.png",
        "images/7.png",
        "images/8.png",
        "images/9.png",
        "images/10.png",
    ]

    for screenshot in screenshots:
        slide_layout = prs.slide_layouts[5]
        slide = prs.slides.add_slide(slide_layout)

        img_path = screenshot
        img = slide.shapes.add_picture(img_path, Inches(0), Inches(0), width=Inches(10), height=Inches(7.5))

        left = int((Inches(10) - img.width) / 2)
        top = int((Inches(7.5) - img.height) / 2)

        img.left = left
        img.top = top

        xml = '''
            <p:transition xmlns:p="http://schemas.openxmlformats.org/presentationml/2006/main" spd="slow" advance="after" advTm="5000">
                <p:dissolve />
            </p:transition>
        '''
        xmlFragment = parse_xml(xml)
        slide.element.insert(-1, xmlFragment)

    prs.save('presentation.pptx')

if __name__ == "__main__":
    create_presentation()
