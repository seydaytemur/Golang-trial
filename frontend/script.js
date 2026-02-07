const API_URL = "http://localhost:8080/api/urunler";
const urunForm = document.getElementById("urunForm");
const urunIdInput = document.getElementById("urunId");
const urunAdiInput = document.getElementById("urunAdi");
const stokMiktariInput = document.getElementById("stokMiktari");
const fiyatInput = document.getElementById("fiyat");
const submitBtn = document.getElementById("submitBtn");
const cancelBtn = document.getElementById("cancelBtn");
const urunListesi = document.getElementById("urunListesi");

document.addEventListener("DOMContentLoaded", fetchUrunler);

async function fetchUrunler() {
    try {
        const response = await fetch(API_URL);
        if (!response.ok) {
            throw new Error(`HTTP Hata! Durum: ${response.status}`);
        }
        const urunler = await response.json();
        renderUrunler(urunler);
    } catch (error) {
        console.error("Ürünler çekilirken bir hata oluştu:", error);
    }
}

function renderUrunler(urunler) {
    urunListesi.innerHTML = '';
    if (urunler && urunler.length > 0) {
        urunler.forEach(urun => {
            const li = document.createElement("li");
            li.innerHTML = `
                <span>${urun.urun_adi} - Stok: ${urun.stok_miktari} - Fiyat: ${urun.fiyat} TL</span>
                <div class="actions">
                    <button class="update-btn" onclick="prepareUpdate(${urun.id}, '${urun.urun_adi}', ${urun.stok_miktari}, ${urun.fiyat})">Güncelle</button>
                    <button class="delete-btn" onclick="deleteUrun(${urun.id})">Sil</button>
                </div>
            `;
            urunListesi.appendChild(li);
        });
    } else {
        urunListesi.innerHTML = '<li>Henüz ürün bulunmamaktadır.</li>';
    }
}

urunForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    const id = urunIdInput.value;
    const urun = {
        urun_adi: urunAdiInput.value,
        stok_miktari: parseInt(stokMiktariInput.value, 10),
        fiyat: parseFloat(fiyatInput.value)
    };

    try {
        if (id) {
            await updateUrun(id, urun);
        } else {
            await createUrun(urun);
        }
    } catch (error) {
        console.error("İşlem sırasında bir hata oluştu:", error);
    }
});

async function createUrun(urun) {
    await fetch(API_URL, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(urun)
    });
    resetForm();
    fetchUrunler();
}

async function deleteUrun(id) {
    const confirmed = confirm("Bu ürünü silmek istediğinizden emin misiniz?");
    if (confirmed) {
        await fetch(`${API_URL}/${id}`, {
            method: "DELETE"
        });
        fetchUrunler();
    }
}

function prepareUpdate(id, ad, stok, fiyat) {
    urunIdInput.value = id;
    urunAdiInput.value = ad;
    stokMiktariInput.value = stok;
    fiyatInput.value = fiyat;
    submitBtn.textContent = "Güncelle";
    cancelBtn.style.display = "inline";
}

async function updateUrun(id, urun) {
    await fetch(`${API_URL}/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(urun)
    });
    resetForm();
    fetchUrunler();
}

cancelBtn.addEventListener("click", resetForm);

function resetForm() {
    urunForm.reset();
    urunIdInput.value = "";
    submitBtn.textContent = "Ürün Ekle";
    cancelBtn.style.display = "none";
}
