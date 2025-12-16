run:
	@air & \
	sleep 2 && \
	if pgrep -x "jprq" > /dev/null; then \
		echo "jprq already running, continue..."; \
	else \
		jprq http 3000; \
	fi