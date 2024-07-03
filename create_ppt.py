from pptx import Presentation
from pptx.util import Inches

def create_presentation():
    prs = Presentation()

    screenshots = [
            'images/weather-today.png',
            'images/weather-tomorrow.png',
            'images/fire.jpg', 
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

    prs.save('presentation.pptx')

if __name__ == "__main__":
    create_presentation()
