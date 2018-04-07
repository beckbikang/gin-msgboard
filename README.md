
a msg board write in gin

install dep

	go get github.com/beckbikang/gin-msgboard
	go get github.com/jinzhu/gorm
	go get github.com/gin-gonic/gin

make sure the mysql config is right in the path conf/config.json

create the table in 

	mesgboard.sql
	
run 

	sh run/start.sh


