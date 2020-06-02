host="cbi-test1"
user="dev"
cbi="/store/app/src/cbi-sh"

ssh -o StrictHostKeyChecking=no ${user}@${host} hostname
ssh ${user}@${host} cat /etc/*-release | grep PRETTY_NAME
ssh ${user}@${host} lscpu | egrep 'Model name|Socket|Thread|NUMA|CPU\(s\)'
ssh ${user}@${host} free -h
ssh ${user}@${host} df -h /store
ssh ${user}@${host} ulimit -aH
echo "================================="
ssh ${user}@${host} mkdir -p ${cbi}
ssh ${user}@${host} ls -la ${cbi}
scp -r [!.]* ${user}@${host}:${cbi}

#ssh dev@${host} "echo 'branch: ${env.BRANCH_NAME}, build: ${buildNumber}' > ${appHome}/${system}-${
#      module
#    }/buildInfo" """
#    sh """ssh dev@${hostname} ${appHome}/${system}-${module}/service/install ${env.BRANCH_NAME}"""
