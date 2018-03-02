--"return": {
--        "code": "0",
--        "data.nickname": "nzlov"
--}
local json = require("json")
function check(rs)
  local r  = json.decode(rs)  
  if (r["code"] ~= "0" )
  then
    return false,"code != 0,is " .. r["code"]
  end
  if (r["data"]["nickname"] ~= "nzlov" )
  then
    return false,"data.nickname != nzlov,is " .. r["data"]["nickname"]
  end
  return true,""
end
