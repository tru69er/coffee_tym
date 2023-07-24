async function doFetch() {
  const req = {
    name: "Cappuccino",
    price: 6.9,
  };

  const f = await fetch("http://127.0.0.1:6969/1", {
    method: "PUT",
    body: JSON.stringify(req),
  });

  console.log(f.statusText);
}

doFetch();
