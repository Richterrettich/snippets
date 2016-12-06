#!/bin/sh

DIRECTORY=testfiles
# -c = create; -J = compress with xz; -f = filename
tar -cJf test.tar.xz $DIRECTORY


# other compression flags are: z = gz; j = bz2
