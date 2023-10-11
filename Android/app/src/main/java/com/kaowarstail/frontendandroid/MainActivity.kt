import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.util.Log
import android.widget.ArrayAdapter
import android.widget.ListView
import com.google.gson.Gson
import com.google.gson.reflect.TypeToken
import okhttp3.*
import java.io.IOException
import com.kaowarstail.frontendandroid.R
import com.kaowarstail.frontendandroid.Product

class MainActivity : AppCompatActivity() {
    private val client = OkHttpClient()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        fetchProducts()
    }

    private fun fetchProducts() {
        val request = Request.Builder()
            .url("http://localhost:8000/items")
            .build()

        client.newCall(request).enqueue(object : Callback {
            override fun onFailure(call: Call, e: IOException) {
                Log.e(TAG, "Error fetching products: ${e.message}")
            }

            override fun onResponse(call: Call, response: Response) {
                val json = response.body?.string()
                val products = Gson().fromJson<List<Product>>(json, object : TypeToken<List<Product>>() {}.type)

                runOnUiThread {
                    // Update UI with fetched products
                    val productList: MutableList<String> = mutableListOf()

                    products.forEach { product ->
                        val productDetails = "Name: ${product.name}\nDescription: ${product.description}\nAvailable Sizes: ${product.sizes}"
                        productList.add(productDetails)
                    }

                    // Utilisez la liste `productList` pour afficher les détails des produits dans votre interface utilisateur
                    // Par exemple, vous pouvez l'afficher dans une ListView ou un TextView
                    val listView: ListView = findViewById(R.id.productListView)
                    val adapter = ArrayAdapter<String>(this@MainActivity, android.R.layout.simple_list_item_1, productList)
                    listView.adapter = adapter
                }


            }
        })
    }

    companion object {
        private const val TAG = "MainActivity"
    }
}
