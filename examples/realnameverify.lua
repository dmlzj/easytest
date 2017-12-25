require("examples.utils");
require("examples.base64");
local img = GetFile("./examples/img.png");
local ba = BASE64Encode(img);
--print(ba);
Req:Param("img",ba);