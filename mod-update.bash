current=$(grep -i "github.com/berserktech/telebot/gh .*" go.mod | awk '{ print $2 }')
latest=$(git rev-parse $@)
echo Updating from $current to $latest
find . -name "*.sum" -type f -delete
find -name *.mod -type f -exec sed -i "s/$current/$latest/g" {} \;
updated=$(grep -i "github.com/berserktech/telebot/gh .*" go.mod | awk '{ print $2 }')
go mod verify
# Let's do it again once we have the latest version
find -name *.mod -type f -exec sed -i "s/$current/$updated/g" {} \;
(cd gh && go mod verify)
(cd gh && go mod verify)
(cd gh && go mod verify)
