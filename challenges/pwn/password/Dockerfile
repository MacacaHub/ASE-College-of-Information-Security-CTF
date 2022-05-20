FROM ubuntu:20.04

RUN apt-get -qq update
RUN apt-get -qqy install xinetd

RUN useradd -m password

RUN chmod 774 /tmp
RUN chmod -R 774 /var/tmp
RUN chmod -R 774 /dev
RUN chmod -R 774 /run
RUN chmod 1733 /tmp /var/tmp /dev/shm
RUN chown -R root:root /home/password

CMD ["/usr/sbin/xinetd","-dontfork"]
