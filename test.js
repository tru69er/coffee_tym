async function doFetch() {
  const req = {
    desc: "3ple espresso shot + hot milk",
    price: 6.9,
  };

  const f = await fetch("http://127.0.0.1:6969/products", {
    method: "POST",
    body: JSON.stringify(req),
  });

  await console.log(await f.text())
}

doFetch();
