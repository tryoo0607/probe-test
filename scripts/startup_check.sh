#!/bin/sh
set -eu

DIR="${PROBE_DIR:-/tmp/probe}"
FILE="$DIR/startup"

[ -f "$FILE" ] || exit 1
val="$(cat "$FILE" 2>/dev/null || echo "")"
[ "$val" = "true" ] && exit 0 || exit 1
