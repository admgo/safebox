#!/bin/bash
set -e

# 检查是否传入参数
if [ -z "$1" ]; then
    echo "用法: $0 <目标目录>"
    exit 1
fi

TARGET_DIR="$1"

# 创建目录（如果不存在）
mkdir -p "$TARGET_DIR/bin"
mkdir -p "$TARGET_DIR/lib"

# 找到系统 Python 位置
PYTHON_PATH=$(command -v python3 || true)

if [ -z "$PYTHON_PATH" ]; then
    echo "系统未找到 python3"
    exit 1
fi

echo "系统 Python 路径: $PYTHON_PATH"

# 复制 python 可执行文件
cp "$PYTHON_PATH" "$TARGET_DIR/bin/python3"

# 复制依赖库
echo "复制依赖库..."
ldd "$PYTHON_PATH" | awk '{ if ($3 ~ /^\//) print $3 }' | while read -r lib; do
    dest="$TARGET_DIR/lib$(dirname $lib)"
    mkdir -p "$dest"
    cp -u "$lib" "$dest/"
done

echo "Python 已复制到 $TARGET_DIR/bin/python3"
echo "依赖库已复制到 $TARGET_DIR/lib"

echo "你可以这样运行:"
echo "  LD_LIBRARY_PATH=$TARGET_DIR/lib $TARGET_DIR/bin/python3 --version"