require("examples.utils");
require("examples.base64");
local img = GetFile("./examples/img.png");
local ba = URLEncode(BASE64Encode(img));
--print(ba);
Req:SendMapString("img="..ba);