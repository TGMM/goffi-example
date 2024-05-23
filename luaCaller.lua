local ffi = require("ffi")
local goLib = ffi.load("./libmain.so")
ffi.cdef[[
int luaEntryPoint();
]]
local retVal = goLib.luaEntryPoint()
print("luaEntryPoint() returned: " .. retVal)