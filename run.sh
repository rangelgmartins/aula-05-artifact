chmod +x ./aula-05-artifact/hello-server.go

./aula-05-artifact/hello-server.go &

sleep 5

for LOGIN in Homer Bart Maggie;
do
echo "$(date): $(curl -s http://localhost:13000/${LOGIN})"
done
