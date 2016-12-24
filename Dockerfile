FROM scratch
ADD build/linux-amd64/govatar-net /
EXPOSE 8000
CMD ["/govatar-net"]
