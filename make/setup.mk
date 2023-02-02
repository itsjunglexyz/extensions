# Copyright (c) 2016 TFG Co <backend@tfgco.com>
# Author: TFG Co <backend@tfgco.com>
#
# Permission is hereby granted, free of charge, to any person obtaining a copy of
# this software and associated documentation files (the "Software"), to deal in
# the Software without restriction, including without limitation the rights to
# use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
# the Software, and to permit persons to whom the Software is furnished to do so,
# subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
# FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
# COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
# IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
# CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

setup:
	# Ensuring librdkafka is installed in Mac OS
	@/bin/bash -c '[ "`uname -s`" == "Darwin" ] && [ "`which brew`" != "" ] && [ ! -d "/usr/local/Cellar/librdkafka" ] && echo "librdkafka was not found. Installing with brew..." && brew install librdkafka; exit 0'
	# Ensuring librdkafka is installed in Debian and Ubuntu
	@/bin/bash -c '[ "`uname -s`" == "Linux" ] && [ "`which apt-get`" != "" ] && echo "Ensuring librdkafka is installed..." && ./debian-install-librdkafka.sh; exit 0'
	@go get -u github.com/onsi/ginkgo/ginkgo
	@go get github.com/gordonklaus/ineffassign
	@go mod tidy

mocks: mongo-mocks dog-mocks pg-mocks redis-mocks kafka-mocks

mongo-mocks:
	@mockgen -source=mongo/interfaces/interfaces.go -destination=mongo/interfaces/mocks.go -package interfaces
	@echo 'created mongo mocks on ./mongo/interfaces/mocks.go'

dog-mocks:
	@mockgen -source=dogstatsd/dogstatsd.go -destination=dogstatsd/mocks/dogstatsd.go -package mocks
	@echo 'created dog mocks on ./dogstatsd/mocks/dogstatsd.go'

pg-mocks:
	@mockgen -source=pg/interfaces/pg.go -destination=pg/mocks/pg.go -package mocks
	@echo 'created pg mocks on ./pg/mocks/pg.go'

redis-mocks:
	@mockgen -source=redis/interfaces/redis.go -destination=redis/mocks/redis.go -package mocks
	@echo 'created redis mocks on ./redis/mocks/redis.go'

kafka-mocks:
	@mockgen -source=kafka/interfaces/kafka.go -destination=kafka/mocks/kafka.go -package mocks
	@echo 'created kafka mocks on ./kafka/mocks/redis.go'
