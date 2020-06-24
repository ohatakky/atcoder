function playatcoder() {
  cd $GOPATH/src/github.com/ohatakky/playgo/atcoder
  t=$(date "+%Y%m%d_%H%M%S")
  d="temp${t}"
  mkdir "${d}"
  cd "${d}"
  cp $GOPATH/src/github.com/ohatakky/playgo/atcoder/main.go main.go
  code .
}

function ac() {
  dir=$GOPATH/src/github.com/ohatakky/playgo/atcoder/
  file=$(ls $dir | tail -n 1)
  code $dir$file
}

playatcoder
