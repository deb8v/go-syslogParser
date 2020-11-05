
1. all launch arguments
1.1 see table 1
    +-------------------------------------------------------+
    |                  all launch arguments                 |
    +-------------------------------------------------------+
    | short  | full Name | Comment                          |
    +--------|-----------|----------------------------------+
    | r      | lockRfc   | RFC3164 RFC3164 RAW              |
    +--------|-----------|----------------------------------+
    | p      | password  | http header password             |
    +--------|-----------|----------------------------------+
    | u      | url       | http collector url               |
    +--------|-----------|----------------------------------+
    | z      | timeZone  | Time Zone, ex. Asia/Novokuznetsk |
    +--------|-----------|----------------------------------+
    | -      | host      | Hostname                         |
    +--------|-----------|----------------------------------+
    | s      | port      | Listen PORT. Default 514         |
    +--------|-----------|----------------------------------+
    | a      | address   | Listen IP. Default 0.0.0.0       |
    +--------|-----------|----------------------------------+
    | d      | debug     | debug messages                   |
    +--------|-----------|----------------------------------+
    | t      | resend    | TCP proxy IP:PORT                |
    +--------|-----------|----------------------------------+
    | g      | journal   | local log file. ex ../logs/p.txt.|
    +-------------------------------------------------------+
                                                    table 1

2. similar description of launch arguments

2.1 input schema
2.1.1 lock input syslog schema
2.1.2 if not present "AUTO"
2.1.3 examples:
2.1.3.1 manage presentation of the programme to select branches of the parser messages
    +------------------------------------------------------+
    |                  We need precision!?                 |
    +------------------------------------------------------+
    | Number | RFC Name | Comment                          |
    +--------|----------|----------------------------------+
    | 2      | RFC3164  | BSD                              |
    +--------|----------|----------------------------------+
    | 1      | RFC5424  |                                  |
    +--------|----------|----------------------------------+
    | 0      | RAW      | third-rate loggers               |
    +--------|----------|----------------------------------+
    | 100    | AUTO     | auto detection                   |
    +------------------------------------------------------+
                                                    table 3
2.1.3.2:
    --lockRfc RFC3164
    -r 100

2.2 collector server:
2.2.1 accepts the incoming log in JSON format for further processing. Using the HTTP\HTTPS standard
2.2.2 if not present "http://sysjournal.(dhcp domain)/get.go:80"
2.2.3 examples:
    --url http://weblog.rt.khai.pw/get.go
    -u http://weblog.rt.khai.pw/get.go

2.3 authorization password
2.3.1 used for authorization using the HTTP header on the server. Len limit 512 chars.
2.3.2 if not present default argument ""
2.3.3 examples:
    --password bartsuha
    -p huiachko228

2.4 collector public name
2.4.1 defined in the http header
2.4.2 if not present "os.Hostname()", if failed "undefined"
2.4.3 examples:
    --host Van Darkholme cluster

2.5 time zone
2.5.1 set ur timeZone offset
2.5.2 if not present 0
2.5.3 examples:
    --timeZone +7
    -z -2




