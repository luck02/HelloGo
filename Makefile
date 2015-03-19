all:	bin/HelloGo
		@echo "Launching at http://localhost:8080/"
			foreman start -p 8080

bin/HelloGo:
		GOBIN=bin go install
