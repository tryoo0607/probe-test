#!/bin/sh
set -eu

DIR="${PROBE_DIR:-/tmp/probe}"
FILE="$DIR/alive"

# 파일이 있고 내용이 "true"면 통과(0), 아니면 실패(1)
[ -f "$FILE" ] || exit 1
val="$(cat "$FILE" 2>/dev/null || echo "")"
[ "$val" = "true" ] && exit 0 || exit 1