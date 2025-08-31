#!/usr/bin/python3
# -*- coding: utf-8 -*-

# @Time: 2025/8/28 11:31
# @Author: Kenley Wang
# @FileName: test.py

import ctypes
import os
import sys
import traceback
# setup sys.excepthook
def excepthook(type, value, tb):
    sys.stderr.write("".join(traceback.format_exception(type, value, tb)))
    sys.stderr.flush()
    sys.exit(-1)

sys.excepthook = excepthook

lib = ctypes.CDLL("./python.so")
lib.DifySeccomp.argtypes = [ctypes.c_uint32, ctypes.c_uint32, ctypes.c_bool]
lib.DifySeccomp.restype = None

lib.DifySeccomp()


code = open("./test.py").read()

exec(code)