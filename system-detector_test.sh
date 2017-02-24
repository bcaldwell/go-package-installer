env GOOS=linux GOARCH=amd64 go build system_detector.go

images=("ubuntu" "centos" "base/archlinux" "alpine" "fedora" "opensuse" "vbatts/slackware" "debian")

for i in "${images[@]}"
do
	output=$(docker run --rm -v `pwd`:/test -w /test ${i} ./system_detector)
	echo "${i}: ${output} \n"
done


