current=$(grep -i "github.com/berserktech/telebot/gh .*" go.mod | awk '{ print $2 }')
latest=$(git rev-parse $@)
echo Updating from $current to $latest
go mod verify
updated=$(grep -i "github.com/berserktech/telebot/gh .*" go.mod | awk '{ print $2 }')
find -name *.mod -type f -exec sed -i "s/$current/$updated/g" {} \;
find . -name "*.sum" -type f -delete
go mod verify
(cd gh && go mod verify)
(cd gh && go mod verify)
(cd gh && go mod verify)
