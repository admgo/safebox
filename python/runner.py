#!/usr/bin/python3
# -*- coding: utf-8 -*-

# @Time: 2025/8/21 13:48
# @Author: Kenley Wang
# @FileName: runner.py


def main(arg1: str, arg2: str) -> dict:
    print(__name__)
    print("__name__")
    return {
        "result": arg1 + arg2,
    }
main(1,3)