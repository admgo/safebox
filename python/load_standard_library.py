#!/usr/bin/env python3
import sys
import shutil
import os


def ignore_pyc_and_cache(dir, files):
    """Ignore __pycache__ and .pyc files"""
    ignored = []
    for f in files:
        if f.endswith(".pyc") or f == "__pycache__":
            ignored.append(f)
    return ignored


def main():
    if len(sys.argv) != 2:
        sys.stderr.write(f"Usage: {sys.argv[0]} <target_dir>\n")
        sys.exit(1)

    target_root = sys.argv[1]
    os.makedirs(target_root, exist_ok=True)

    try:
        # skip sys.path[0] (current working dir)
        for p in sys.path[1:]:
            if not p or not os.path.exists(p):
                continue

            dest = os.path.join(target_root, p.lstrip(os.sep))
            os.makedirs(os.path.dirname(dest), exist_ok=True)

            if os.path.isdir(p):
                if os.path.exists(dest):
                    shutil.rmtree(dest)
                shutil.copytree(p, dest, symlinks=True, ignore=ignore_pyc_and_cache)
            else:
                shutil.copy2(p, dest)

    except Exception as e:
        sys.stderr.write(f"Error: {e}\n")
        sys.exit(1)


if __name__ == "__main__":
    main()