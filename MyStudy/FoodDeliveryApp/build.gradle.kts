plugins {
    kotlin("jvm") version "1.9.22" apply false
    kotlin("android") version "1.9.22" apply false
    kotlin("kapt") version "1.9.22" apply false
}

buildscript {
    dependencies {
        classpath("com.android.tools.build:gradle:8.1.0")
        classpath("org.jetbrains.kotlin:kotlin-gradle-plugin:1.8.22")
    }
    repositories {
        google()
        mavenCentral()
    }
}

allprojects {
    repositories {
        google()
        mavenCentral()
    }
}
