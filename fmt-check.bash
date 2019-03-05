if [[ $(gofmt -l .) ]]; then exit 1; fi
if [[ $(cd gh && gofmt -l .) ]]; then exit 1; fi
if [[ $(cd tg && gofmt -l .) ]]; then exit 1; fi
