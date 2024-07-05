from pptx import Presentation
from pptx.util import Inches
from datetime import datetime, timedelta
from pptx.oxml import parse_xml
from pptx.opc.constants import RELATIONSHIP_TYPE as RT
from lxml import etree
import warnings
import os

def set_pres_repeat(prs):
    prs_part = prs.part
    prs_props_part = prs_part.part_related_by(RT.PRES_PROPS)
    presentationPr = parse_xml(prs_props_part.blob)

    p = "http://schemas.openxmlformats.org/presentationml/2006/main"

    showPr_elements = presentationPr.findall(f'.//{{{p}}}showPr')
    if not showPr_elements:
        showPr_element = etree.Element(f'{{{p}}}showPr', loop="true", restart="always")
        presentationPr.append(showPr_element)
    else:
        for showPr in showPr_elements:
            showPr.set("loop", "true")
            showPr.set("restart", "always")

    prs_props_part._blob = etree.tostring(presentationPr)

def hide_last_slide(prs):
    slides = prs.slides
    if len(slides) > 0:
        last_slide = slides[-1]
        last_slide.element.set('show', '0')
        print("Last slide hidden successfully.")
    else:
        print("No slides to hide.")

def create_presentation():
    warnings.filterwarnings("ignore", message="Duplicate name:")
    presentation_path = 'presentation.pptx'
    
    if os.path.exists(presentation_path):
        prs = Presentation(presentation_path)

        xml_slides = prs.slides._sldIdLst  
        slide_ids = [slide_id for slide_id in xml_slides]

        for slide_id in slide_ids:
            xml_slides.remove(slide_id)
    else:
        prs = Presentation()
    
    set_pres_repeat(prs)

    today = datetime.now()
    tomorrow = today + timedelta(days=1)

    screenshots = [
        f'images/{today.strftime("%d-%m-%Y")}.png',
        f'images/{tomorrow.strftime("%d-%m-%Y")}.png',
        f'images/{today.strftime("%d-%m-%Y")}-fire.jpg',
        'images/placeholder.png'
    ]

    transition_xml_template = '''
        <p:transition xmlns:p="http://schemas.openxmlformats.org/presentationml/2006/main" spd="slow" advance="after" advTm="1000">
            <p:wipe />
        </p:transition>
    '''

    for i, screenshot in enumerate(screenshots):
        if not os.path.exists(screenshot):
            print(f"File not found: {screenshot}. Skipping this screenshot.")
            continue

        print(screenshot)
        slide_layout = prs.slide_layouts[5]
        slide = prs.slides.add_slide(slide_layout)

        img_path = screenshot
        try:
            img = slide.shapes.add_picture(img_path, Inches(0), Inches(0), width=Inches(10), height=Inches(7.5))

            left = int((Inches(10) - img.width) / 2)
            top = int((Inches(7.5) - img.height) / 2)

            img.left = left
            img.top = top
        except Exception as e:
            print(f"Error adding picture {img_path}: {e}")
            continue

        transition_xml = transition_xml_template
        transition_fragment = parse_xml(transition_xml)

        slide.element.append(transition_fragment)

        print(f"Transition applied to slide {i+1}")

    hide_last_slide(prs)
    prs.save(presentation_path)

if __name__ == "__main__":
    create_presentation()
