host="mac.local"
user="yykhomenko"
appHome="/Users/${user}/app/src/github.com/cbi-sh/files"

ssh -o StrictHostKeyChecking=no ${user}@${host} \
hostname; \
cat /etc/*-release | grep PRETTY_NAME; \
lscpu | egrep 'Model name|Socket|Thread|NUMA|CPU\(s\)'; \
free -h; \
df -h /; \
ulimit -aH; \
echo 'appHome: '${appHome};
#echo "================================="
#ssh ${user}@${host} mkdir -p ${cbi}
#ssh ${user}@${host} ls -la ${cbi}
#scp -r [!.]* ${user}@${host}:${cbi}

#ssh dev@${host} "echo 'branch: ${env.BRANCH_NAME}, build: ${buildNumber}' > ${appHome}/${system}-${
#      module
#    }/buildInfo" """
#    sh """ssh dev@${hostname} ${appHome}/${system}-${module}/service/install ${env.BRANCH_NAME}"""
