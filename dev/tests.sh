#-- PARSER Checker
logger --server 192.168.88.238 -d -P514 --rfc3164 --priority 0 "-" -S 2048 -t "t1"
logger --server 192.168.88.238 -d -P514 --rfc3164 --priority warning "ebal" -S 2048 -t "t2"
logger --server 192.168.88.238 -d -P514 --rfc3164 --priority info "ebal<22>;TABLE DROM TEST;" -S 2048 -t "t2"

logger --server 192.168.88.238 -d -P514 --rfc5424 --priority 0 "-" -S 2048 -t "t3"
logger --server 192.168.88.238 -d -P514 --rfc5424 --priority warning "ebal" -S 2048 -t "t3"
logger --server 192.168.88.238 -d -P514 --rfc5424 --priority info "ebal<22>;TABLE DROM TEST;" -S 2048 -t "t3"



#-- RFC autoChecker
logger --server 192.168.88.238 -d -P514 --rfc3164 --priority 0 "-" -S 2048 -t "t1"
logger --server 192.168.88.238 -d -P514 --rfc5424 --priority 0 "-" -S 2048 -t "t1"

logger --server 192.168.88.238 -d -P514 --rfc3164 --priority warning "ebal" -S 2048 -t "t2"
logger --server 192.168.88.238 -d -P514 --rfc5424 --priority warning "ebal" -S 2048 -t "t2"

logger --server 192.168.88.238 -d -P514 --rfc3164 --priority info "ebal<22>;TABLE DROM TEST;" -S 2048 -t "t3"
logger --server 192.168.88.238 -d -P514 --rfc5424 --priority info "ebal<22>;TABLE DROM TEST;" -S 2048 -t "t3"

logger --server 192.168.88.238 -d -P514 --rfc5424 --priority warning "ebal" -S 2048 -t "t3"
logger --server 192.168.88.238 -d -P514 --rfc3164 --priority warning "ebal" -S 2048 -t "t2"
logger --server 192.168.88.238 -d -P514 "ebal" -S 2048