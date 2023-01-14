FROM busybox:stable-musl
ENV PORT=8000
LABEL maintainer="John Toniutti <john@jotone.eu>"

ADD www /www/

EXPOSE $PORT

HEALTHCHECK CMD nc -z localhost $PORT

# Create a basic webserver and run it until the container is stopped
CMD echo "httpd started" && trap "exit 0;" TERM INT; httpd -v -p $PORT -h /www -f & wait
