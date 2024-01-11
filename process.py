import json

custom_domain = "https://emotion.sky123.top"
custom_file = "onetwo"
custom_name = "一二"
size = 332
page = 5

pageSize = size // page

includeItems = []
totalCount = 1
for pageCount in range(1, page + 1):
    dictData = {}
    dictData["name"] = f"{custom_name}第{pageCount}弹"
    dictData["type"] = "image"
    dictData["items"] = []
    for itemCount in range(1, pageSize + 1):
        item = {}
        item["key"] = f"{custom_file}-{totalCount}"
        item["val"] = f"{custom_domain}/{custom_file}/data/{totalCount}.gif"
        dictData["items"].append(item)
        totalCount = totalCount + 1
    fileName = f"{custom_file}-{pageCount}.json"
    includeItems.append(f"{custom_domain}/{custom_file}/{fileName}")
    with open("./" + custom_file + "/" + fileName, "wb") as f:
        f.write(json.dumps(dictData, indent=4, ensure_ascii=False).encode())

with open("./" + custom_file + "/include.json", "wb") as f:
    f.write(json.dumps(includeItems, indent=4, ensure_ascii=False).encode())

print("--------处理完成--------")
