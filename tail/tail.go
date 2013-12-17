//       -c, --bytes=K
//              output the last K bytes; alternatively,  use  -c  +K  to  output
//              bytes starting with the Kth of each file
//
//       -f, --follow[={name|descriptor}]
//              output appended data as the file grows; -f, --follow, and --fol‐
//              low=descriptor are equivalent
//
//       -F     same as --follow=name --retry
//
//       -n, --lines=K
//              output the last K lines, instead of the last 10; or use -n +K to
//              output lines starting with the Kth
//
//       --max-unchanged-stats=N
//              with  --follow=name,  reopen  a  FILE which has not changed size
//              after N (default 5) iterations to see if it has been unlinked or
//              renamed  (this  is  the  usual case of rotated log files).  With
//              inotify, this option is rarely useful.
//
//       --pid=PID
//              with -f, terminate after process ID, PID dies
//
//       -q, --quiet, --silent
//              never output headers giving file names
//
//       --retry
//              keep trying to open a file even when it is or becomes inaccessi‐
//              ble; useful when following by name, i.e., with --follow=name
//
//       -s, --sleep-interval=N
//              with -f, sleep for approximately N seconds (default 1.0) between
//              iterations.  With inotify and --pid=P, check process P at  least
//              once every N seconds.
//
//       -v, --verbose
//              always output headers giving file names
//
//       --help display this help and exit
//
//       --version
//              output version information and exit

package main


