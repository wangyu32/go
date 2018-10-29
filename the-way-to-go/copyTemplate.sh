#!bin/bash
#目的：在任意位置执行got命令，能够在当前目录下生成一个以template.go为模板的go文件。
#      且文件命名规则满足一下：
#      a.当前目录下没有temp.go文件，则文件名为temp.go
#      b.当前目录下已经有temp.go文件，则以"temp" + "时间戳".go为文件名，如temp1540836197.go
#

tempPath=/go/the-way-to-go/
tempFile=template.go
tempFileAbs=$tempPath$tempFile
#echo $tempFileAbs

currPath=`pwd`
#echo $currPath
fileName=$currPath/$tempFile

#echo $fileName

fileExist=0
test -e $fileName && fileExist=1

#echo $fileExist

if [ $fileExist == 1 ]; then
	t=`date "+%s"`
#	echo $t
	fileName=$currPath/"template"$t".go"
#	echo $fileName
fi

cp $tempFileAbs $fileName
echo "cp $tempFileAbs $fileName OK"
