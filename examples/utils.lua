--读取指定文件  
function GetFile(filename)  
    local file = io.open(filename, "r")
    local data = file:read("*all")  
    file:close()  
    return data   
end 

function URLEncode(s)  
    s = string.gsub(s, "([^%w%.%- ])", function(c) return string.format("%%%02X", string.byte(c)) end)  
   return string.gsub(s, " ", "+")  
end  
 
function URLDecode(s)  
   s = string.gsub(s, '%%(%x%x)', function(h) return string.char(tonumber(h, 16)) end)  
   return s  
end