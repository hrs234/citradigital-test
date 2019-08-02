FROM alpine
RUN mkdir /app
WORKDIR /app
COPY solve-3 /app
CMD ["./solve-3"]
EXPOSE 5000
