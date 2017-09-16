# Mudio

This service wraps the Sox binary, providing a limited networked interface. It has a single endpoint which takes the location of an MP3 file (probably in S3), downloads the file, generates a bunch of modified files, and then uploads those to S3.