import os
import logging
import pathlib
import json
from fastapi import FastAPI, Form, HTTPException
from fastapi.responses import FileResponse
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()
logger = logging.getLogger("uvicorn")
logger.level = logging.INFO
images = pathlib.Path(__file__).parent.resolve() / "images"
origins = [os.environ.get("FRONT_URL", "http://localhost:3000")]
json_path = "sample.json"
app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=False,
    allow_methods=["GET", "POST", "PUT", "DELETE"],
    allow_headers=["*"],
)


@app.get("/")
def root():
    return {"message": "Hello, world!"}

@app.get("/items")
def return_sample_json():
    sample_json_load = json.load(open(json_path,'r'))
    print(sample_json_load)
    return sample_json_load



@app.post("/items")
def add_item(name: str = Form(),category:str = Form()):
    logger.info(f"Receive item: {name}")
    logger.info(f"Recive category : {category}")
    item_dict = {"name":name,"category":category}
    with open("./edited.json", "w") as js:
        json.dump(item_dict, js, indent = 4)
        print("success creating json")
    return {"message": f"item received: name = {name} category ={category}"}


@app.get("/image/{image_name}")
async def get_image(image_name):
    # Create image path
    image = images / image_name

    if not image_name.endswith(".jpg"):
        raise HTTPException(status_code=400, detail="Image path does not end with .jpg")

    if not image.exists():
        logger.debug(f"Image not found: {image}")
        image = images / "default.jpg"

    return FileResponse(image)
