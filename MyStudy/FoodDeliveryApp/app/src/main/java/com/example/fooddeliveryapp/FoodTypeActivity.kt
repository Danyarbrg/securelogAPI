package com.example.foodapp

import android.os.Bundle
import android.view.MenuItem
import android.widget.TextView
import androidx.appcompat.app.AppCompatActivity
import com.example.foodapp.R
import android.widget.ImageView
import androidx.appcompat.widget.Toolbar
import androidx.lifecycle.lifecycleScope
import com.bumptech.glide.Glide
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory


class FoodTypeActivity : AppCompatActivity() {

    private lateinit var tvName: TextView
    private lateinit var ivImage: ImageView
    private lateinit var tvDescription: TextView

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_food_type)

        // 1) Toolbar как ActionBar
        val toolbar = findViewById<Toolbar>(R.id.toolbar)
        setSupportActionBar(toolbar)
        supportActionBar?.setDisplayHomeAsUpEnabled(true)

        // 2) view binding через findViewById
        tvName = findViewById(R.id.tv_food_type)
        ivImage = findViewById(R.id.iv_food_image)
        tvDescription = findViewById(R.id.tv_description)

        // 3) Отображаем название
        val foodType = intent.getStringExtra("food_type") ?: ""
        tvName.text = foodType
        supportActionBar?.title = foodType

        // 4) Загружаем данные о блюде из TheMealDB
        lifecycleScope.launch {
            try {
                val api = Retrofit.Builder()
                    .baseUrl("https://www.themealdb.com/")
                    .addConverterFactory(GsonConverterFactory.create())
                    .build()
                    .create(MealApi::class.java)

                val response = withContext(Dispatchers.IO) {
                    api.searchMealsByName(foodType)
                }

                // берём первое блюдо из списка
                val meal = response.meals?.firstOrNull()
                meal?.let {
                    // загрузка картинки через Glide (не забудь добавить зависимость!)
                    Glide.with(this@FoodTypeActivity)
                        .load(it.strMealThumb)
                        .into(ivImage)

                    tvDescription.text = it.strInstructions
                }
            } catch (e: Exception) {
                tvDescription.text = "Не удалось загрузить данные"
            }
        }
    }

    // Обработка клика по стрелке «назад»
    override fun onOptionsItemSelected(item: MenuItem): Boolean {
        if (item.itemId == android.R.id.home) {
            finish()
            return true
        }
        return super.onOptionsItemSelected(item)
    }
}
