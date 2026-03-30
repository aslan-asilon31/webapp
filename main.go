package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>AslanAsilon Go — DummyJSON API Manager</title>
  <link href="https://fonts.googleapis.com/css2?family=Syne:wght@400;600;700;800&family=DM+Sans:wght@300;400;500&display=swap" rel="stylesheet"/>
  <style>
    :root {
      --white:    #ffffff;
      --off:      #f4f6fb;
      --ocean1:   #e8f0fe;
      --ocean2:   #c7d9fc;
      --blue1:    #4a7fe5;
      --blue2:    #2456c8;
      --blue3:    #0f2d72;
      --purple1:  #7c5cbf;
      --purple2:  #a78bfa;
      --purple3:  #5b21b6;
      --teal:     #06b6d4;
      --shadow:   0 4px 32px rgba(36,86,200,.10);
      --radius:   18px;
    }

    *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

    body {
      font-family: 'DM Sans', sans-serif;
      background: var(--white);
      color: #1a1a2e;
      min-height: 100vh;
      overflow-x: hidden;
    }

    /* ── BACKGROUND MESH ── */
    body::before {
      content: '';
      position: fixed; inset: 0; z-index: -1;
      background:
        radial-gradient(ellipse 80% 60% at 10% 0%, #dde9ff 0%, transparent 55%),
        radial-gradient(ellipse 60% 50% at 90% 10%, #ede3ff 0%, transparent 50%),
        radial-gradient(ellipse 50% 70% at 50% 100%, #c7d9fc 0%, transparent 60%),
        var(--white);
    }

    /* ── HEADER ── */
    header {
      display: flex; align-items: center; justify-content: space-between;
      padding: 22px 48px;
      border-bottom: 1px solid rgba(74,127,229,.12);
      backdrop-filter: blur(12px);
      position: sticky; top: 0; z-index: 100;
      background: rgba(255,255,255,.82);
    }
    .logo {
      font-family: 'Syne', sans-serif;
      font-size: 1.35rem; font-weight: 800; letter-spacing: -.5px;
      background: linear-gradient(120deg, var(--blue2), var(--purple1));
      -webkit-background-clip: text; -webkit-text-fill-color: transparent;
    }
    .logo span { color: var(--teal); -webkit-text-fill-color: var(--teal); }
    .badge {
      background: linear-gradient(135deg, var(--blue1), var(--purple1));
      color: #fff; font-size: .72rem; font-weight: 600;
      padding: 4px 12px; border-radius: 99px; letter-spacing: .5px;
    }

    /* ── HERO ── */
    .hero {
      padding: 64px 48px 40px;
      max-width: 1200px; margin: 0 auto;
    }
    .hero h1 {
      font-family: 'Syne', sans-serif;
      font-size: clamp(2rem, 4vw, 3.2rem);
      font-weight: 800; line-height: 1.1;
      background: linear-gradient(130deg, var(--blue3) 20%, var(--purple1) 70%, var(--teal) 100%);
      -webkit-background-clip: text; -webkit-text-fill-color: transparent;
    }
    .hero p { margin-top: 12px; color: #5a6a8a; font-size: 1.05rem; max-width: 560px; }

    /* ── ENDPOINT SELECTOR ── */
    .controls {
      max-width: 1200px; margin: 0 auto;
      padding: 0 48px 24px;
      display: flex; gap: 12px; flex-wrap: wrap; align-items: center;
    }
    .ep-btn {
      border: 1.5px solid var(--ocean2);
      background: var(--white);
      border-radius: 10px;
      padding: 8px 20px;
      font-family: 'DM Sans', sans-serif;
      font-size: .88rem; font-weight: 500;
      cursor: pointer; transition: all .2s;
      color: var(--blue2);
    }
    .ep-btn:hover, .ep-btn.active {
      background: linear-gradient(135deg, var(--blue1), var(--purple1));
      border-color: transparent; color: #fff;
      box-shadow: 0 4px 16px rgba(74,127,229,.35);
    }

    /* ── MAIN GRID ── */
    .main-grid {
      max-width: 1200px; margin: 0 auto;
      padding: 0 48px 64px;
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 28px;
    }
    @media(max-width: 768px) { .main-grid { grid-template-columns: 1fr; } .hero, .controls, .main-grid { padding-left: 20px; padding-right: 20px; } header { padding: 16px 20px; } }

    /* ── CARD ── */
    .card {
      background: rgba(255,255,255,.85);
      border: 1px solid rgba(74,127,229,.13);
      border-radius: var(--radius);
      padding: 28px;
      box-shadow: var(--shadow);
      backdrop-filter: blur(8px);
      transition: box-shadow .25s;
    }
    .card:hover { box-shadow: 0 8px 40px rgba(36,86,200,.15); }
    .card-title {
      font-family: 'Syne', sans-serif;
      font-size: 1rem; font-weight: 700;
      color: var(--blue3);
      margin-bottom: 18px;
      display: flex; align-items: center; gap: 8px;
    }
    .card-title .dot {
      width: 8px; height: 8px; border-radius: 50%;
      background: linear-gradient(135deg, var(--blue1), var(--purple1));
    }

    /* ── FORM ── */
    label { display: block; font-size: .8rem; font-weight: 500; color: #6478a0; margin-bottom: 5px; }
    input, select, textarea {
      width: 100%; padding: 10px 14px;
      border: 1.5px solid var(--ocean2);
      border-radius: 10px;
      font-family: 'DM Sans', sans-serif; font-size: .9rem;
      color: #1a1a2e; background: var(--off);
      transition: border-color .2s, box-shadow .2s;
      outline: none;
    }
    input:focus, select:focus, textarea:focus {
      border-color: var(--blue1);
      box-shadow: 0 0 0 3px rgba(74,127,229,.15);
      background: #fff;
    }
    textarea { resize: vertical; min-height: 100px; }
    .field { margin-bottom: 14px; }
    .row2 { display: grid; grid-template-columns: 1fr 1fr; gap: 12px; }

    /* ── BUTTONS ── */
    .btn {
      padding: 11px 26px; border: none; border-radius: 10px;
      font-family: 'DM Sans', sans-serif; font-size: .9rem; font-weight: 500;
      cursor: pointer; transition: all .2s;
    }
    .btn-primary {
      background: linear-gradient(135deg, var(--blue1), var(--purple1));
      color: #fff; box-shadow: 0 4px 14px rgba(74,127,229,.3);
    }
    .btn-primary:hover { transform: translateY(-1px); box-shadow: 0 6px 20px rgba(74,127,229,.4); }
    .btn-outline {
      background: transparent; border: 1.5px solid var(--ocean2); color: var(--blue2);
    }
    .btn-outline:hover { border-color: var(--blue1); background: var(--ocean1); }
    .btn-danger { background: linear-gradient(135deg, #e05c5c, #b91c1c); color: #fff; }
    .btn-actions { display: flex; gap: 10px; margin-top: 4px; }

    /* ── RESPONSE BOX ── */
    .response-box {
      background: #f0f4ff;
      border: 1.5px solid var(--ocean2);
      border-radius: 12px;
      padding: 16px;
      font-family: 'Courier New', monospace;
      font-size: .8rem;
      color: #2456c8;
      min-height: 120px;
      max-height: 320px;
      overflow-y: auto;
      white-space: pre-wrap;
      word-break: break-word;
      line-height: 1.6;
    }
    .response-label {
      font-size: .75rem; font-weight: 600; color: #8a9cc0;
      text-transform: uppercase; letter-spacing: .8px;
      margin-bottom: 8px;
    }
    .status-ok  { color: #16a34a; font-weight: 700; }
    .status-err { color: #dc2626; font-weight: 700; }

    /* ── PRODUCT LIST ── */
    .product-item {
      display: flex; align-items: center; gap: 14px;
      padding: 12px 0; border-bottom: 1px solid var(--ocean1);
    }
    .product-item:last-child { border-bottom: none; }
    .product-thumb {
      width: 48px; height: 48px; object-fit: cover;
      border-radius: 10px; background: var(--ocean1);
      flex-shrink: 0;
    }
    .product-info { flex: 1; }
    .product-name { font-weight: 600; font-size: .9rem; color: #1a1a2e; }
    .product-meta { font-size: .78rem; color: #7a8aaa; margin-top: 2px; }
    .product-price {
      font-family: 'Syne', sans-serif; font-weight: 700;
      color: var(--blue2); font-size: .95rem;
    }

    /* ── LOADING SPINNER ── */
    .spinner {
      display: inline-block; width: 18px; height: 18px;
      border: 2.5px solid rgba(255,255,255,.3);
      border-top-color: #fff;
      border-radius: 50%; animation: spin .7s linear infinite;
      vertical-align: middle; margin-right: 6px;
    }
    @keyframes spin { to { transform: rotate(360deg); } }

    /* ── STATS ── */
    .stats-row { display: grid; grid-template-columns: repeat(3, 1fr); gap: 14px; margin-bottom: 20px; }
    .stat-card {
      background: linear-gradient(135deg, var(--ocean1), var(--ocean2));
      border-radius: 12px; padding: 16px 14px; text-align: center;
    }
    .stat-num {
      font-family: 'Syne', sans-serif; font-size: 1.5rem;
      font-weight: 800; color: var(--blue2);
    }
    .stat-lbl { font-size: .75rem; color: #6478a0; margin-top: 2px; }

    /* ── METHOD BADGE ── */
    .method {
      font-size: .7rem; font-weight: 700; padding: 2px 8px;
      border-radius: 6px; letter-spacing: .5px; margin-right: 6px;
    }
    .GET    { background: #dcfce7; color: #15803d; }
    .POST   { background: #dbeafe; color: #1d4ed8; }
    .PUT    { background: #fef9c3; color: #a16207; }
    .DELETE { background: #fee2e2; color: #b91c1c; }

    /* ── TABS ── */
    .tabs { display: flex; gap: 4px; margin-bottom: 18px; }
    .tab {
      padding: 7px 18px; border-radius: 8px; font-size: .85rem; font-weight: 500;
      cursor: pointer; border: none; background: transparent; color: #8a9cc0;
      transition: all .2s;
    }
    .tab.active { background: var(--ocean1); color: var(--blue2); font-weight: 600; }
    .tab:hover:not(.active) { background: var(--off); }

    .hidden { display: none !important; }

    /* ── FOOTER ── */
    footer {
      text-align: center; padding: 24px;
      font-size: .78rem; color: #a0aec0;
      border-top: 1px solid var(--ocean1);
    }
    footer span {
      background: linear-gradient(90deg, var(--blue1), var(--purple1));
      -webkit-background-clip: text; -webkit-text-fill-color: transparent; font-weight: 700;
    }
  </style>
</head>
<body>

<header>
  <div class="logo">AslanAsilon<span>Go</span></div>
  <div style="display:flex;align-items:center;gap:12px">
    <span style="font-size:.8rem;color:#8a9cc0">DummyJSON API Manager</span>
    <span class="badge">v1.0</span>
  </div>
</header>

<div class="hero">
  <h1>Marketplace API<br>Management Console</h1>
  <p>Kelola produk, cart, dan user melalui DummyJSON REST API dengan antarmuka yang elegan.</p>
</div>

<!-- Endpoint selector -->
<div class="controls">
  <button class="ep-btn active" onclick="setEndpoint('products')">🛍 Products</button>
  <button class="ep-btn" onclick="setEndpoint('carts')">🛒 Carts</button>
  <button class="ep-btn" onclick="setEndpoint('users')">👤 Users</button>
  <button class="ep-btn" onclick="setEndpoint('auth')">🔑 Auth</button>
</div>

<div class="main-grid">

  <!-- LEFT: Actions -->
  <div>
    <!-- PRODUCTS -->
    <div id="panel-products">
      <div class="card" style="margin-bottom:20px">
        <div class="card-title"><div class="dot"></div>Fetch Products</div>
        <div class="tabs">
          <button class="tab active" onclick="switchTab(this,'pt-list')">List</button>
          <button class="tab" onclick="switchTab(this,'pt-single')">Single</button>
          <button class="tab" onclick="switchTab(this,'pt-search')">Search</button>
          <button class="tab" onclick="switchTab(this,'pt-cat')">Category</button>
        </div>
        <div id="pt-list">
          <div class="row2">
            <div class="field"><label>Limit</label><input id="p-limit" type="number" value="5" min="1" max="30"/></div>
            <div class="field"><label>Skip</label><input id="p-skip" type="number" value="0" min="0"/></div>
          </div>
          <div class="btn-actions">
            <button class="btn btn-primary" onclick="fetchProducts()"><span id="sp1"></span>GET All</button>
          </div>
        </div>
        <div id="pt-single" class="hidden">
          <div class="field"><label>Product ID</label><input id="p-id" type="number" value="1" min="1"/></div>
          <div class="btn-actions">
            <button class="btn btn-primary" onclick="fetchProductById()"><span id="sp2"></span>GET by ID</button>
          </div>
        </div>
        <div id="pt-search" class="hidden">
          <div class="field"><label>Query</label><input id="p-q" type="text" placeholder="e.g. phone"/></div>
          <div class="btn-actions">
            <button class="btn btn-primary" onclick="searchProducts()"><span id="sp3"></span>Search</button>
          </div>
        </div>
        <div id="pt-cat" class="hidden">
          <div class="field">
            <label>Category</label>
            <select id="p-cat">
              <option value="smartphones">Smartphones</option>
              <option value="laptops">Laptops</option>
              <option value="fragrances">Fragrances</option>
              <option value="skincare">Skincare</option>
              <option value="groceries">Groceries</option>
              <option value="home-decoration">Home Decoration</option>
            </select>
          </div>
          <div class="btn-actions">
            <button class="btn btn-primary" onclick="fetchByCategory()"><span id="sp4"></span>GET Category</button>
          </div>
        </div>
      </div>

      <div class="card">
        <div class="card-title"><div class="dot"></div>Add / Update / Delete Product</div>
        <div class="tabs">
          <button class="tab active" onclick="switchTab(this,'pm-add')">Add</button>
          <button class="tab" onclick="switchTab(this,'pm-update')">Update</button>
          <button class="tab" onclick="switchTab(this,'pm-delete')">Delete</button>
        </div>
        <div id="pm-add">
          <div class="field"><label>Title</label><input id="new-title" placeholder="Product name"/></div>
          <div class="row2">
            <div class="field"><label>Price ($)</label><input id="new-price" type="number" placeholder="99.99"/></div>
            <div class="field"><label>Stock</label><input id="new-stock" type="number" placeholder="50"/></div>
          </div>
          <div class="field"><label>Description</label><textarea id="new-desc" placeholder="Product description..."></textarea></div>
          <div class="btn-actions">
            <button class="btn btn-primary" onclick="addProduct()">POST Add</button>
          </div>
        </div>
        <div id="pm-update" class="hidden">
          <div class="field"><label>Product ID</label><input id="upd-id" type="number" value="1"/></div>
          <div class="field"><label>New Title</label><input id="upd-title" placeholder="Updated title"/></div>
          <div class="field"><label>New Price ($)</label><input id="upd-price" type="number" placeholder=""/></div>
          <div class="btn-actions">
            <button class="btn btn-primary" onclick="updateProduct()">PUT Update</button>
          </div>
        </div>
        <div id="pm-delete" class="hidden">
          <div class="field"><label>Product ID to Delete</label><input id="del-id" type="number" value="1"/></div>
          <div class="btn-actions">
            <button class="btn btn-danger" onclick="deleteProduct()">DELETE Remove</button>
          </div>
        </div>
      </div>
    </div>

    <!-- CARTS -->
    <div id="panel-carts" class="hidden">
      <div class="card" style="margin-bottom:20px">
        <div class="card-title"><div class="dot"></div>Fetch Carts</div>
        <div class="tabs">
          <button class="tab active" onclick="switchTab(this,'ct-all')">All Carts</button>
          <button class="tab" onclick="switchTab(this,'ct-user')">User's Cart</button>
        </div>
        <div id="ct-all">
          <div class="btn-actions" style="margin-top:4px">
            <button class="btn btn-primary" onclick="apiCall('GET','https://dummyjson.com/carts?limit=5')">GET Carts</button>
          </div>
        </div>
        <div id="ct-user" class="hidden">
          <div class="field"><label>User ID</label><input id="cart-uid" type="number" value="5"/></div>
          <div class="btn-actions">
            <button class="btn btn-primary" onclick="apiCall('GET','https://dummyjson.com/carts/user/'+document.getElementById('cart-uid').value)">GET User Cart</button>
          </div>
        </div>
      </div>
      <div class="card">
        <div class="card-title"><div class="dot"></div>Add to Cart</div>
        <div class="field"><label>User ID</label><input id="c-uid" type="number" value="1"/></div>
        <div class="field"><label>Product ID</label><input id="c-pid" type="number" value="1"/></div>
        <div class="field"><label>Quantity</label><input id="c-qty" type="number" value="1"/></div>
        <div class="btn-actions">
          <button class="btn btn-primary" onclick="addToCart()">POST Add Cart</button>
        </div>
      </div>
    </div>

    <!-- USERS -->
    <div id="panel-users" class="hidden">
      <div class="card" style="margin-bottom:20px">
        <div class="card-title"><div class="dot"></div>Fetch Users</div>
        <div class="row2">
          <div class="field"><label>Limit</label><input id="u-limit" type="number" value="5"/></div>
          <div class="field"><label>Skip</label><input id="u-skip" type="number" value="0"/></div>
        </div>
        <div class="btn-actions">
          <button class="btn btn-primary" onclick="fetchUsers()">GET Users</button>
          <button class="btn btn-outline" onclick="apiCall('GET','https://dummyjson.com/users/1')">GET User #1</button>
        </div>
      </div>
      <div class="card">
        <div class="card-title"><div class="dot"></div>Search Users</div>
        <div class="field"><label>Query</label><input id="u-q" placeholder="e.g. Emily"/></div>
        <div class="btn-actions">
          <button class="btn btn-primary" onclick="apiCall('GET','https://dummyjson.com/users/search?q='+encodeURIComponent(document.getElementById('u-q').value))">Search</button>
        </div>
      </div>
    </div>

    <!-- AUTH -->
    <div id="panel-auth" class="hidden">
      <div class="card">
        <div class="card-title"><div class="dot"></div>Login</div>
        <div class="field"><label>Username</label><input id="auth-user" value="emilys"/></div>
        <div class="field"><label>Password</label><input id="auth-pass" type="password" value="emilyspass"/></div>
        <div class="field">
          <label>Expire in (minutes)</label>
          <input id="auth-exp" type="number" value="30"/>
        </div>
        <div class="btn-actions">
          <button class="btn btn-primary" onclick="doLogin()">POST Login</button>
        </div>
        <p style="margin-top:12px;font-size:.78rem;color:#8a9cc0">Demo: emilys / emilyspass</p>
      </div>
    </div>
  </div>

  <!-- RIGHT: Response + Preview -->
  <div>
    <div class="card" style="margin-bottom:20px">
      <div class="card-title"><div class="dot"></div>API Response</div>
      <div style="display:flex;justify-content:space-between;align-items:center;margin-bottom:8px">
        <div id="resp-status" style="font-size:.82rem"></div>
        <button class="btn btn-outline" style="padding:5px 12px;font-size:.78rem" onclick="clearResp()">Clear</button>
      </div>
      <div class="response-label">JSON Output</div>
      <div class="response-box" id="resp-box">// Response will appear here...</div>
    </div>

    <div class="card" id="preview-card">
      <div class="card-title"><div class="dot"></div>Product Preview</div>
      <div class="stats-row">
        <div class="stat-card"><div class="stat-num" id="s-total">—</div><div class="stat-lbl">Total</div></div>
        <div class="stat-card"><div class="stat-num" id="s-skip">—</div><div class="stat-lbl">Skip</div></div>
        <div class="stat-card"><div class="stat-num" id="s-limit">—</div><div class="stat-lbl">Limit</div></div>
      </div>
      <div id="product-list"></div>
    </div>
  </div>
</div>

<footer>
  Built with <span>AslanAsilon Go</span> · Powered by DummyJSON API · Port :65534
</footer>

<script>
  const BASE = 'https://dummyjson.com';
  let currentEndpoint = 'products';

  function setEndpoint(ep) {
    currentEndpoint = ep;
    document.querySelectorAll('.ep-btn').forEach(b => b.classList.remove('active'));
    event.target.classList.add('active');
    ['products','carts','users','auth'].forEach(p => {
      document.getElementById('panel-'+p).classList.toggle('hidden', p !== ep);
    });
  }

  function switchTab(btn, targetId) {
    const parent = btn.closest('.card');
    parent.querySelectorAll('.tab').forEach(t => t.classList.remove('active'));
    btn.classList.add('active');
    // hide siblings
    const panels = parent.querySelectorAll('[id]');
    panels.forEach(p => {
      if (p.id === targetId) p.classList.remove('hidden');
      else if (p.id.startsWith(targetId.split('-')[0]+'-')) p.classList.add('hidden');
    });
  }

  function setStatus(code, ms) {
    const el = document.getElementById('resp-status');
    const ok = code >= 200 && code < 300;
    el.innerHTML = '<span class="'+(ok?'status-ok':'status-err')+'">'+code+'</span> · '+ms+'ms · '+(ok?'✓ OK':'✗ Error');
  }

  async function apiCall(method, url, body) {
    const box = document.getElementById('resp-box');
    box.textContent = 'Loading...';
    const t = Date.now();
    try {
      const opts = { method, headers: {'Content-Type':'application/json'} };
      if (body) opts.body = JSON.stringify(body);
      const r = await fetch(url, opts);
      const data = await r.json();
      setStatus(r.status, Date.now()-t);
      box.textContent = JSON.stringify(data, null, 2);
      return data;
    } catch(e) {
      box.textContent = 'Error: ' + e.message;
    }
  }

  function clearResp() {
    document.getElementById('resp-box').textContent = '// Response cleared.';
    document.getElementById('resp-status').innerHTML = '';
  }

  async function fetchProducts() {
    const limit = document.getElementById('p-limit').value;
    const skip  = document.getElementById('p-skip').value;
    const data  = await apiCall('GET', BASE+'/products?limit='+limit+'&skip='+skip);
    if (data) renderProducts(data);
  }

  async function fetchProductById() {
    const id = document.getElementById('p-id').value;
    await apiCall('GET', BASE+'/products/'+id);
  }

  async function searchProducts() {
    const q = document.getElementById('p-q').value;
    const data = await apiCall('GET', BASE+'/products/search?q='+encodeURIComponent(q));
    if (data) renderProducts(data);
  }

  async function fetchByCategory() {
    const cat = document.getElementById('p-cat').value;
    const data = await apiCall('GET', BASE+'/products/category/'+cat);
    if (data) renderProducts(data);
  }

  function renderProducts(data) {
    document.getElementById('s-total').textContent = data.total ?? '—';
    document.getElementById('s-skip').textContent  = data.skip  ?? '—';
    document.getElementById('s-limit').textContent = data.limit ?? '—';
    const list = document.getElementById('product-list');
    list.innerHTML = '';
    (data.products || (data.id ? [data] : [])).slice(0,5).forEach(p => {
      list.innerHTML += '<div class="product-item">'+
        '<img class="product-thumb" src="'+(p.thumbnail||'')+('" onerror="this.src=\'\'")+'>' +
        '<div class="product-info">'+
          '<div class="product-name">'+p.title+'</div>'+
          '<div class="product-meta">'+p.category+' · ⭐ '+(p.rating||'N/A')+'</div>'+
        '</div>'+
        '<div class="product-price">$'+p.price+'</div>'+
      '</div>';
    });
  }

  async function addProduct() {
    const body = {
      title: document.getElementById('new-title').value,
      price: parseFloat(document.getElementById('new-price').value),
      stock: parseInt(document.getElementById('new-stock').value),
      description: document.getElementById('new-desc').value,
    };
    await apiCall('POST', BASE+'/products/add', body);
  }

  async function updateProduct() {
    const id = document.getElementById('upd-id').value;
    const body = {};
    const t = document.getElementById('upd-title').value;
    const p = document.getElementById('upd-price').value;
    if (t) body.title = t;
    if (p) body.price = parseFloat(p);
    await apiCall('PUT', BASE+'/products/'+id, body);
  }

  async function deleteProduct() {
    const id = document.getElementById('del-id').value;
    await apiCall('DELETE', BASE+'/products/'+id);
  }

  async function addToCart() {
    const body = {
      userId: parseInt(document.getElementById('c-uid').value),
      products: [{ id: parseInt(document.getElementById('c-pid').value), quantity: parseInt(document.getElementById('c-qty').value) }]
    };
    await apiCall('POST', BASE+'/carts/add', body);
  }

  async function fetchUsers() {
    const limit = document.getElementById('u-limit').value;
    const skip  = document.getElementById('u-skip').value;
    await apiCall('GET', BASE+'/users?limit='+limit+'&skip='+skip);
  }

  async function doLogin() {
    const body = {
      username: document.getElementById('auth-user').value,
      password: document.getElementById('auth-pass').value,
      expiresInMins: parseInt(document.getElementById('auth-exp').value),
    };
    await apiCall('POST', BASE+'/auth/login', body);
  }

  // auto-load sample products on startup
  window.addEventListener('load', fetchProducts);
</script>
</body>
</html>`)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":65534", nil)
}
