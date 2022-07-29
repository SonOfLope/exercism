git add $1/* 
git commit -m "$1 exercise completed heavy_check_mark"
rm -rf $1
echo "$1/" >> .gitignore

