async function doFetch() {
  const req = {
    name: "Cappuccino",
    desc: "3ple espresso shot + hot milk",
    price: 6.9,
  };

  const f = await fetch("http://127.0.0.1:6969/2", {
    method: "DELETE",
  });

  console.log(await f.text())
}

doFetch();
