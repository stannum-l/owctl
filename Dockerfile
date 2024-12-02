ARG FROM=ubuntu:latest
FROM $FROM

FROM scratch
COPY --from=build /out /
ENTRYPOINT ["owctl"]